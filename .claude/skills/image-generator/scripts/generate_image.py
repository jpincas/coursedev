#!/usr/bin/env python3
"""
Generate images using OpenAI's image generation API for training course modules.
"""
import os
import sys
import argparse
import json
import requests
from pathlib import Path

try:
    from openai import OpenAI
except ImportError:
    print("ERROR: openai package not installed. Run: pip install openai", file=sys.stderr)
    sys.exit(1)


def generate_image(prompt, output_path=None, size="1536x1024", model="gpt-image-1.5", quality="high"):
    """
    Generate an image using OpenAI's API.

    Args:
        prompt: Text description of the image to generate
        output_path: Where to save the image (optional)
        size: Image dimensions (gpt-image-1.5: 1024x1024/1024x1536/1536x1024/auto; dall-e-3: 1024x1024/1792x1024/1024x1792)
        model: Model to use (default: gpt-image-1.5)
        quality: Image quality (gpt-image-1.5: low/medium/high/auto; dall-e-3: standard/hd)

    Returns:
        Path to saved image or URL if output_path not specified
    """
    api_key = os.getenv("OPENAI_API_KEY")
    if not api_key:
        raise ValueError("OPENAI_API_KEY environment variable not set")

    client = OpenAI(api_key=api_key)

    print(f"Generating image with prompt: {prompt[:100]}...", file=sys.stderr)
    print(f"Model: {model}, Size: {size}, Quality: {quality}", file=sys.stderr)

    try:
        response = client.images.generate(
            model=model,
            prompt=prompt,
            n=1,
            size=size,
            quality=quality
        )

        # Handle different response formats (URL vs base64)
        image_url = None
        if hasattr(response.data[0], 'url') and response.data[0].url:
            image_url = response.data[0].url
        elif hasattr(response.data[0], 'b64_json') and response.data[0].b64_json:
            # Image returned as base64, not URL
            import base64
            image_data = base64.b64decode(response.data[0].b64_json)
            if output_path:
                output_file = Path(output_path)
                output_file.parent.mkdir(parents=True, exist_ok=True)
                with open(output_file, 'wb') as f:
                    f.write(image_data)
                print(f"Image saved to: {output_path}", file=sys.stderr)
                return str(output_file.absolute())
            else:
                raise ValueError("Image returned as base64 but no output path specified")
        else:
            raise ValueError(f"Unexpected response format. Attributes: {dir(response.data[0])}")

        print(f"Image generated successfully: {image_url}", file=sys.stderr)

        if output_path:
            # Download and save the image
            print(f"Downloading image to {output_path}...", file=sys.stderr)
            image_response = requests.get(image_url)
            image_response.raise_for_status()

            output_file = Path(output_path)
            output_file.parent.mkdir(parents=True, exist_ok=True)

            with open(output_file, 'wb') as f:
                f.write(image_response.content)

            print(f"Image saved to: {output_path}", file=sys.stderr)
            return str(output_file.absolute())
        else:
            return image_url

    except Exception as e:
        print(f"ERROR: Failed to generate image: {str(e)}", file=sys.stderr)
        raise


def main():
    parser = argparse.ArgumentParser(
        description="Generate images for training course modules using OpenAI API"
    )
    parser.add_argument(
        "prompt",
        help="Text description of the image to generate"
    )
    parser.add_argument(
        "--output", "-o",
        help="Path to save the generated image (optional, returns URL if not specified)"
    )
    parser.add_argument(
        "--size", "-s",
        default="1536x1024",
        choices=["1024x1024", "1024x1536", "1536x1024", "auto", "1792x1024", "1024x1792"],
        help="Image size (gpt-image-1.5: 1024x1024/1024x1536/1536x1024/auto; dall-e-3: 1024x1024/1792x1024/1024x1792) (default: 1536x1024)"
    )
    parser.add_argument(
        "--model", "-m",
        default="gpt-image-1.5",
        help="Model to use (default: gpt-image-1.5)"
    )
    parser.add_argument(
        "--quality", "-q",
        default="high",
        choices=["low", "medium", "high", "auto", "standard", "hd"],
        help="Image quality (gpt-image-1.5: low/medium/high/auto; dall-e-3: standard/hd) (default: high)"
    )
    parser.add_argument(
        "--json",
        action="store_true",
        help="Output result as JSON"
    )

    args = parser.parse_args()

    try:
        result = generate_image(
            prompt=args.prompt,
            output_path=args.output,
            size=args.size,
            model=args.model,
            quality=args.quality
        )

        if args.json:
            print(json.dumps({
                "success": True,
                "path": result if args.output else None,
                "url": result if not args.output else None
            }))
        else:
            print(result)

        return 0

    except Exception as e:
        if args.json:
            print(json.dumps({
                "success": False,
                "error": str(e)
            }))
        else:
            print(f"ERROR: {str(e)}", file=sys.stderr)
        return 1


if __name__ == "__main__":
    sys.exit(main())
